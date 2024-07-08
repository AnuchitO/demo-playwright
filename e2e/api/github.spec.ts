import {test, expect} from '@playwright/test'

const REPO = 'test-repo-1'
const USER = 'github-username'

test('should create a bug report', async ({request}) => {
  const newIssue = await request.post(`/repos/${USER}/${REPO}/issues`, {
    data: {
      title: '[Bug] report 1',
      body: 'Bug description',
    },
  })

  expect(newIssue.statusText()).toContain('OK')
  expect(newIssue.url()).toContain(
    'https://api.github.com/repos/github-username/test-repo-1/issues'
  )
  expect(newIssue.ok()).toBeTruthy()

  const issues = await request.get(`/repos/${USER}/${REPO}/issues`)
  expect(issues.ok()).toBeTruthy()
  expect(await issues.json()).toContainEqual(
    expect.objectContaining({
      title: '[Bug] report 1',
      body: 'Bug description',
    })
  )
})
