import { Component } from '@angular/core'
import { NgIf, AsyncPipe, JsonPipe } from '@angular/common'
import { Store, select } from '@ngrx/store'
import { Observable } from 'rxjs'
import { SettingsFeature } from 'store/settings'
import { BaseCardComponent } from 'components/cards/base-card/base-card.component'
import { NewsCardComponent } from 'components/cards/news-card/news-card.component'
import { TimeCardComponent } from 'components/cards/time-card/time-card.component'
import { WeatherCardComponent } from 'components/cards/weather-card/weather-card.component'
import { animations } from './home.animations'

@Component({
    standalone: true,
    selector: 'cmp-home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.scss'],
    animations: [animations],
    imports: [
        NgIf,
        AsyncPipe,
        JsonPipe,
        BaseCardComponent,
        NewsCardComponent,
        TimeCardComponent,
        WeatherCardComponent,
    ],
})
export default class HomeComponent {
    public loaded$: Observable<boolean> = this.store.pipe(select(SettingsFeature.selectLoaded))

    public constructor(private readonly store: Store) {}
}
