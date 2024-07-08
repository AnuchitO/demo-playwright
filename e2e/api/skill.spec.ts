import {test, expect} from '@playwright/test'
import {client} from '../database/postgres'

test.beforeAll(async () => {
  try {
    await client.connect()
    await client.query(
      'INSERT INTO skill (key, name, description, logo, levels, tags) VALUES ($1, $2, $3, $4, $5, $6);',
      [
        'kotlin',
        'Kotlin',
        'Kotlin is a cross-platform...',
        'https://logo.com/kotlin',
        '[{"name": "Beginner", "level": 1, "descriptions": ["basic knowledge ..."]},{"name": "Intermediate", "level": 2, "descriptions": ["complex programs..."]}]',
        ['kotlin', 'android'],
      ]
    )
  } catch (error) {
    console.error('Error connecting to database:', error)
  }
})

test.afterAll(async () => {
  try {
    await client.query("DELETE FROM skill where key='kotlin';")
    await client.end()
  } catch (error) {
    console.error('Error connecting to database:', error)
  }
})

test('should response one skill when request /skills/:key', async ({
  request,
}) => {
  const reps = await request.get('/skills/kotlin')

  expect(reps.ok()).toBeTruthy()
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: 'success',
      data: {
        key: 'kotlin',
        name: 'Kotlin',
        description: expect.any(String),
        logo: expect.any(String),
        levels: expect.arrayContaining([
          {
            key: expect.any(String),
            name: 'Beginner',
            brief: expect.any(String),
            descriptions: expect.arrayContaining([expect.any(String)]),
            level: 1,
          },
          {
            key: expect.any(String),
            name: 'Intermediate',
            brief: expect.any(String),
            descriptions: expect.arrayContaining([expect.any(String)]),
            level: 2,
          },
        ]),
        tags: expect.arrayContaining(['kotlin', 'android']),
      },
    })
  )
})
