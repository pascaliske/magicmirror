import { Component, inject } from '@angular/core'
import { AsyncPipe } from '@angular/common'
import { Store, select } from '@ngrx/store'
import { Observable } from 'rxjs'
import { SettingsFeature } from 'store/settings'
import { NewsCardComponent } from 'components/cards/news-card/news-card.component'
import { TimeCardComponent } from 'components/cards/time-card/time-card.component'
import { WeatherCardComponent } from 'components/cards/weather-card/weather-card.component'
import { animations } from './home.animations'

@Component({
    selector: 'cmp-home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.scss'],
    animations: [animations],
    imports: [AsyncPipe, NewsCardComponent, TimeCardComponent, WeatherCardComponent],
})
export default class HomeComponent {
    private readonly store: Store = inject(Store)

    public loaded$: Observable<boolean> = this.store.pipe(select(SettingsFeature.selectLoaded))
}
