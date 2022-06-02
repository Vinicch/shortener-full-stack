export type ShortenerError = {
  ERR_CODE: string
  Description: string
}

export type UrlInfo = {
  alias: string
  original: string
  shortened: string
  elapsedTime: string
}

export type MostVisitedUrl = {
  url: string
  short_url: string
  visits: number
}
