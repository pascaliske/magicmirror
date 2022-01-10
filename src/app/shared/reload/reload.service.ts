import { Injectable } from '@angular/core'
import { timer } from 'rxjs'
import { take } from 'rxjs/operators'

@Injectable({
    providedIn: 'root',
})
export class ReloadService {
    /**
     * Reloads the browser window.
     *
     * @param delay Optional delay in milliseconds.
     */
    public reload(delay?: number): void {
        // reload directly
        if (!delay || delay === 0) {
            window.location.reload()
            return
        }

        // reload after delay
        timer(delay)
            .pipe(take(1))
            .subscribe(() => {
                window.location.reload()
            })
    }
}
