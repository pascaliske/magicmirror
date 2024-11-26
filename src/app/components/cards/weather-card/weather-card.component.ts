import { Component, ChangeDetectionStrategy } from '@angular/core'
import { NgIf, NgFor, AsyncPipe } from '@angular/common'
import { HttpParams } from '@angular/common/http'
import { Observable, EMPTY } from 'rxjs'
import { concatMap, map } from 'rxjs/operators'
import { Cacheable } from 'ts-cacheable'
import { environment } from 'environments/environment'
import { Settings } from 'store/settings'
import { FromTimestampPipe } from 'shared/datetime/from-timestamp.pipe'
import { ToLocalDatePipe } from 'shared/datetime/to-local-date.pipe'
import { ToLocalTimePipe } from 'shared/datetime/to-local-time.pipe'
import { ToDayPipe } from 'shared/datetime/to-day.pipe'
import { ToDayNamePipe } from 'shared/datetime/to-day-name.pipe'
import { ToWeekPipe } from 'shared/datetime/to-week.pipe'
import { ToMonthPipe } from 'shared/datetime/to-month.pipe'
import { ToMonthNamePipe } from 'shared/datetime/to-month-name.pipe'
import { ToYearPipe } from 'shared/datetime/to-year.pipe'
import { ToHumanNumberPipe } from 'shared/numbers/to-human.pipe'
import { BaseCardComponent, repeatAfter } from '../base-card/base-card.component'
import { WeatherData } from './weather-data'

@Component({
    selector: 'cmp-weather-card',
    templateUrl: './weather-card.component.html',
    styleUrls: ['./weather-card.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [
        NgIf,
        NgFor,
        AsyncPipe,
        FromTimestampPipe,
        ToLocalDatePipe,
        ToLocalTimePipe,
        ToDayPipe,
        ToDayNamePipe,
        ToWeekPipe,
        ToMonthPipe,
        ToMonthNamePipe,
        ToYearPipe,
        ToHumanNumberPipe,
    ],
})
export class WeatherCardComponent extends BaseCardComponent {
    private readonly data$: Observable<WeatherData> = this.settings$.pipe(
        repeatAfter(),
        concatMap(({ data }) => this.fetchWeatherData(data)),
    )

    public current$: Observable<WeatherData['current']> = this.data$.pipe(
        map(({ current }) => {
            // convert visibility to km
            current.visibility /= 1000
            return current
        }),
    )

    public daily$: Observable<WeatherData['daily']> = this.data$.pipe(
        map(({ daily }) => daily.slice(0, 8)),
    )

    @Cacheable()
    private fetchWeatherData(settings: Settings): Observable<WeatherData> {
        if (!settings?.apiKeys?.openWeather) {
            return EMPTY
        }

        const url: string = environment.production
            ? 'https://api.openweathermap.org/data/2.5/onecall'
            : '/assets/mockups/weather.json'
        const params: HttpParams = new HttpParams()
            .set('appid', settings?.apiKeys?.openWeather)
            .set('exclude', 'minutely')
            .set('lat', settings?.location?.latitude ?? '')
            .set('lon', settings?.location?.longitude ?? '')
            .set('units', 'metric')
            .set('lang', this.locale)

        return this.http.get<WeatherData>(url, { params })
    }
}
