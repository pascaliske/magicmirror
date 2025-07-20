import { HttpClient } from '@angular/common/http'
import { TranslateLoader } from '@ngx-translate/core'
import { SentryService } from '@pascaliske/ngx-sentry'
import { Observable, Subject, EMPTY } from 'rxjs'
import { catchError } from 'rxjs/operators'
import { Cacheable, LocalStorageStrategy } from 'ts-cacheable'

const cacheBuster$: Subject<void> = new Subject<void>()

export class TranslationLoader implements TranslateLoader {
    public constructor(
        private readonly http: HttpClient,
        private readonly sentry: SentryService,
    ) {}

    /**
     * Fetches the translations from the API.
     *
     * @returns - An observable resolving to an object of translations.
     */
    @Cacheable({
        storageStrategy: LocalStorageStrategy,
        cacheBusterObserver: cacheBuster$,
        maxAge: 3600000,
    })
    public getTranslation(language: string): Observable<Record<string, string>> {
        return this.http.get<Record<string, string>>(`/assets/translations/${language}.json`).pipe(
            catchError((error: Error) => {
                cacheBuster$.next()

                this.sentry.withScope(scope => {
                    scope.setExtra('error', error)
                    this.sentry.captureException(new Error(`Couldn't load translation.`))
                })

                return EMPTY
            }),
        )
    }
}
