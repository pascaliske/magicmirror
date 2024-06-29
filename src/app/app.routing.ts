import type { RouterFeatures, Routes } from '@angular/router'
import { withDisabledInitialNavigation, withInMemoryScrolling } from '@angular/router'

export const features: RouterFeatures[] = [
    withDisabledInitialNavigation(),
    withInMemoryScrolling({
        anchorScrolling: 'enabled',
        scrollPositionRestoration: 'enabled',
    }),
]

export const routes: Routes = [
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'home',
    },
    {
        path: 'home',
        loadComponent: () => import('./pages/home/home.component'),
    },
    {
        path: '**',
        redirectTo: 'home',
    },
]
