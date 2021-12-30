import { NgModule } from '@angular/core'
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome'
import { SharedModule } from 'shared/shared.module'
import { IconComponent } from './icon.component'

@NgModule({
    imports: [FontAwesomeModule, SharedModule],
    declarations: [IconComponent],
    exports: [IconComponent],
})
export class IconModule {}
