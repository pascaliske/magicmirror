import { Component } from '@angular/core'
import { Store, select } from '@ngrx/store'
import { Observable } from 'rxjs'
import { AppState } from 'store'
import { SettingsState } from 'store/settings'
import { animations } from './home.animations'

@Component({
    selector: 'cmp-home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.scss'],
    animations: [animations],
})
export class HomeComponent {
    public settings$: Observable<SettingsState> = this.store.pipe(select('settings'))

    public constructor(private readonly store: Store<AppState>) {}
}
