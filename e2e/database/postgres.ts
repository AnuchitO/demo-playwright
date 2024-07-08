import {Client} from 'pg'

export const client = new Client({
  user: 'postgres',
  password: 'postgres',
  host: '127.0.0.1',
  port: 5432,
  database: 'app',
})

export async function select() {
  try {
    await client.connect()
    console.log('Connected to Postgres database')

    // Execute a raw SQL query
    const result = await client.query('SELECT * FROM skill;')
    console.log('result:', result.rows)

    await client.end()
  } catch (error) {
    console.error('Error connecting to database:', error)
  }
}
