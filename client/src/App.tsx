import { Button, Card, CardContent, Stack, TextField, Typography } from '@mui/material'
import { Container } from '@mui/system'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { MostVisitedUrl, UrlInfo, ShortenerError } from './dtos'

const baseUrl = 'http://localhost:8080'

function App() {
  const [url, setUrl] = useState('')
  const [alias, setAlias] = useState('')
  const [error, setError] = useState<ShortenerError | null>(null)
  const [urlInfo, setUrlInfo] = useState<UrlInfo | null>(null)
  const [urls, setUrls] = useState<MostVisitedUrl[]>([])

  const shortenUrl = async () => {
    try {
      const response = await axios.post<UrlInfo>(`${baseUrl}/create?url=${url}&CUSTOM_ALIAS=${alias}`)

      setError(null)
      setUrlInfo(response.data)
    } catch (e: any) {
      setUrlInfo(null)
      setError(e.response.data)
    }
  }

  const getMostVisited = async () => {
    try {
      const response = await axios.get<MostVisitedUrl[]>(`${baseUrl}/most-visited`)

      setError(null)
      setUrls(response.data)
    } catch (e: any) {
      setUrls([])
      setError(e.response.data)
    }
  }

  useEffect(() => {
    getMostVisited()
  }, [])

  return (
    <Container>
      <Stack spacing={2}>
        <TextField required label='URL' variant='standard' value={url} onChange={(e) => setUrl(e.target.value)} />
        <TextField label='Custom Alias' variant='standard' value={alias} onChange={(e) => setAlias(e.target.value)} />

        <Button variant='contained' onClick={shortenUrl}>
          Shorten URL
        </Button>
      </Stack>

      <br />

      {error && (
        <Card>
          <CardContent>
            <Typography>
              <b>Error {error.ERR_CODE}:</b>
              <br />
              {error.Description}
            </Typography>
          </CardContent>
        </Card>
      )}

      {urlInfo && (
        <Card>
          <CardContent>
            <Typography align='center'>Result</Typography>
            <Typography>
              <b>Alias: </b> {urlInfo.alias}
              <br />
              <b>Shortened URL: </b>
              <a href={urlInfo.shortened} target='_blank' rel='noreferrer'>
                {urlInfo.shortened}
              </a>
            </Typography>
          </CardContent>
        </Card>
      )}

      <br />

      <Card>
        <CardContent>
          <Button variant='outlined' onClick={getMostVisited}>
            Refresh
          </Button>
          <Typography align='center'>Most visited URLs</Typography>
          <ul>
            {urls.map((url) => (
              <li key={url.short_url}>
                <Typography>
                  <a href={url.short_url} target='_blank' rel='noreferrer'>
                    {url.short_url}
                  </a>{' '}
                  ({url.url}) - {url.visits} visits
                </Typography>
              </li>
            ))}
          </ul>
        </CardContent>
      </Card>
    </Container>
  )
}

export default App
