interface WeatherDataBase {
    dt: number
    clouds: number
    humidity: number
    pressure: number
    wind_speed: number
    wind_deg: number
    wind_gust: number
    weather: { id: number; main: string; description: string }[]
}

interface WeatherDataTempSingle {
    temp: number
    feels_like: number
}

interface WeatherDataTempMulti {
    temp: { min: number; max: number; day: number; night: number; eve: number; morn: number }
    feels_like: { day: number; night: number; eve: number; morn: number }
}

interface WeatherDataVisibility {
    visibility: number
}

interface WeatherDataRain {
    rain: number
}

interface WeatherDataSun {
    sunrise: number
    sunset: number
}

interface WeatherDataMoon {
    moonrise: number
    moonset: number
    moon_phase: number
}

interface WeatherDataAlert {
    sender_name: string
    event: string
    start: number
    end: number
    description: string
    tags: string[]
}

export type WeatherDataCurrent = WeatherDataBase &
    WeatherDataTempSingle &
    WeatherDataVisibility &
    WeatherDataRain &
    WeatherDataSun
export type WeatherDataHourly = WeatherDataBase &
    WeatherDataTempSingle &
    WeatherDataVisibility &
    WeatherDataRain
export type WeatherDataDaily = WeatherDataBase &
    WeatherDataTempMulti &
    WeatherDataRain &
    WeatherDataSun &
    WeatherDataMoon

export interface WeatherData {
    current: WeatherDataCurrent
    hourly: WeatherDataHourly[]
    daily: WeatherDataDaily[]
    alerts: WeatherDataAlert[]
}
