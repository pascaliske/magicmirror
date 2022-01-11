import { NgModule } from '@angular/core'
import { SharedModule } from 'shared/shared.module'
import { IconModule } from '../icon/icon.module'
import { BaseCardComponent } from './base-card/base-card.component'
import { TimeCardComponent } from './time-card/time-card.component'

@NgModule({
    imports: [SharedModule, IconModule],
    declarations: [BaseCardComponent, TimeCardComponent],
    exports: [BaseCardComponent, TimeCardComponent],
})
export class CardsModule {}
