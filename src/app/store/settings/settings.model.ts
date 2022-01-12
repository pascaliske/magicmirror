export interface Settings {
    language: string
    units: 'metric' | 'imperial'
    location: {
        latitude: number
        longitude: number
    }
    feeds: string[]
    apiKeys: {
        openWeather: string
    }
}
