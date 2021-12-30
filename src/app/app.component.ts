import { Component, OnInit } from '@angular/core'
import { Store } from '@ngrx/store'
import { Observable } from 'rxjs'
import { AppState } from 'store'
import { SettingsState, LoadSettings } from 'store/settings'
import { animations } from './app.animations'

@Component({
    selector: 'cmp-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
    animations,
})
export class AppComponent implements OnInit {
    public settings$: Observable<SettingsState> = this.store.select('settings')

    public constructor(private readonly store: Store<AppState>) {}

    public ngOnInit(): void {
        this.store.dispatch(LoadSettings())
    }
}
