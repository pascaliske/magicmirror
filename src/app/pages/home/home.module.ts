import { NgModule } from '@angular/core'
import { RouterModule, Routes } from '@angular/router'
import { SharedModule } from 'shared/shared.module'
import { CardsModule } from 'components/cards/cards.module'
import { HomeComponent } from './home.component'

export const routes: Routes = [
    {
        path: '',
        component: HomeComponent,
    },
]

@NgModule({
    imports: [RouterModule.forChild(routes), SharedModule, CardsModule],
    declarations: [HomeComponent],
})
export class HomeModule {}
