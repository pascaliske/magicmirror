export interface Settings {
    location: {
        latitude: number
        longitude: number
    }
    feeds: string[]
    apiKeys: {
        openWeather: string
    }
}
