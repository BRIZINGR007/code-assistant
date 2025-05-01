import { Routes } from '@angular/router';
import { authGuard } from './core/guards/auth.guard';

export const routes: Routes = [
    {
        path: 'login',
        loadComponent: () =>
            import('./layout/auth/login/login.component').then(m => m.LoginComponent),
    },
    {
        path: 'signup',
        loadComponent: () =>
            import('./layout/auth/signup/signup.component').then(m => m.SignupComponent),
    },
    {
        path: '',
        loadComponent: () =>
            import('./layout/dashboard/dashboard.component').then(m => m.DashboardComponent),
        canActivate: [authGuard],
    },
    {
        path: "chat-session",
        loadComponent: () =>
            import('./layout/chat/chat.component').then(m => m.ChatComponent),
        canActivate: [authGuard],

    },
];

