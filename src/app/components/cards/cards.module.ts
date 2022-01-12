import { NgModule } from '@angular/core'
import { SharedModule } from 'shared/shared.module'
import { IconModule } from '../icon/icon.module'
import { BaseCardComponent } from './base-card/base-card.component'
import { NewsCardComponent } from './news-card/news-card.component'
import { TimeCardComponent } from './time-card/time-card.component'
import { WeatherCardComponent } from './weather-card/weather-card.component'

@NgModule({
    imports: [SharedModule, IconModule],
    declarations: [BaseCardComponent, NewsCardComponent, TimeCardComponent, WeatherCardComponent],
    exports: [BaseCardComponent, NewsCardComponent, TimeCardComponent, WeatherCardComponent],
})
export class CardsModule {}
