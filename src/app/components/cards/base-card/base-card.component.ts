import { Component, ChangeDetectionStrategy } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Store } from '@ngrx/store'
import { Observable } from 'rxjs'
import { AppState } from 'store'
import { SettingsState } from 'store/settings'

@Component({
    selector: 'cmp-base-card',
    templateUrl: './base-card.component.html',
    styleUrls: ['./base-card.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class BaseCardComponent {
    protected settings$: Observable<SettingsState> = this.store.select('settings')

    public constructor(protected http: HttpClient, private readonly store: Store<AppState>) {}
}
