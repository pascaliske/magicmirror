import { Component, OnInit, DestroyRef, inject } from '@angular/core'
import { takeUntilDestroyed } from '@angular/core/rxjs-interop'
import { Router, RouterOutlet } from '@angular/router'
import { Store } from '@ngrx/store'
import { NgProgressbar } from 'ngx-progressbar'
import { NgProgressHttp } from 'ngx-progressbar/http'
import { NgProgressRouter } from 'ngx-progressbar/router'
import { SocketService, SocketAction } from 'shared/socket/socket.service'
import { HealthService, Status } from 'shared/health/health.service'
import { ReloadService } from 'shared/reload/reload.service'
import { SettingsActions } from 'store/settings'

@Component({
    selector: 'cmp-root',
    templateUrl: './app.component.html',
    imports: [RouterOutlet, NgProgressbar, NgProgressHttp, NgProgressRouter],
})
export class AppComponent implements OnInit {
    private readonly destroy: DestroyRef = inject(DestroyRef)

    private readonly router: Router = inject(Router)

    private readonly store: Store = inject(Store)

    private readonly socket: SocketService = inject(SocketService)

    private readonly health: HealthService = inject(HealthService)

    private readonly reload: ReloadService = inject(ReloadService)

    public ngOnInit(): void {
        // delayed initial navigation
        setTimeout(() => this.router.initialNavigation())

        // fetch settings from server
        this.store.dispatch(SettingsActions.load())

        // watch for reload actions
        this.socket
            .subscribe(SocketAction.Reload)
            .pipe(takeUntilDestroyed(this.destroy))
            .subscribe(() => {
                this.reload.reload()
            })

        // watch for unhealthy state
        this.health
            .watch(Status.Unavailable)
            .pipe(takeUntilDestroyed(this.destroy))
            .subscribe(() => {
                this.reload.reload()
            })
    }
}
