import { Component, OnInit } from '@angular/core'
import { UntilDestroy, untilDestroyed } from '@ngneat/until-destroy'
import { Store } from '@ngrx/store'
import { Observable } from 'rxjs'
import { HealthService, Status } from 'shared/health/health.service'
import { ReloadService } from 'shared/reload/reload.service'
import { AppState } from 'store'
import { SettingsState, LoadSettings } from 'store/settings'
import { animations } from './app.animations'

@UntilDestroy()
@Component({
    selector: 'cmp-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
    animations,
})
export class AppComponent implements OnInit {
    public settings$: Observable<SettingsState> = this.store.select('settings')

    public constructor(
        private readonly store: Store<AppState>,
        private readonly health: HealthService,
        private readonly reload: ReloadService,
    ) {}

    public ngOnInit(): void {
        this.store.dispatch(LoadSettings())

        this.health
            .watch(Status.Unavailable)
            .pipe(untilDestroyed(this))
            .subscribe(() => {
                this.reload.reload()
            })
    }
}
