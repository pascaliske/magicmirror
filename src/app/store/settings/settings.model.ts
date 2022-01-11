export interface Settings {
    language: string
    units: 'metric' | 'imperial'
    location: {
        latitude: number
        longitude: number
    }
    apiKeys: {
        openWeather: string
    }
}
