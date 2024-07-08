import {test, expect} from '@playwright/test'

test('should response one skill when request /skills/:key', async ({
  request,
}) => {
  const reps = await request.get('/skills/go')

  expect(reps.ok()).toBeTruthy()
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      status: 'success',
      data: {
        key: 'go',
        name: 'Go',
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
        tags: expect.arrayContaining(['go', 'golang']),
      },
    })
  )
})
