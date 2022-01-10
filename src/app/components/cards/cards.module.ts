import { NgModule } from '@angular/core'
import { SharedModule } from 'shared/shared.module'
import { IconModule } from '../icon/icon.module'
import { BaseCardComponent } from './base-card/base-card.component'
import { TimeCardComponent } from './time-card/time-card.component'
import { WeatherCardComponent } from './weather-card/weather-card.component'

@NgModule({
    imports: [SharedModule, IconModule],
    declarations: [BaseCardComponent, TimeCardComponent, WeatherCardComponent],
    exports: [BaseCardComponent, TimeCardComponent, WeatherCardComponent],
})
export class CardsModule {}
