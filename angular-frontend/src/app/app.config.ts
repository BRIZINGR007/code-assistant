import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideMarkdown, MarkedOptions, MarkedRenderer, MARKED_OPTIONS } from 'ngx-markdown';

import { routes } from './app.routes';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { provideToastr } from 'ngx-toastr';
import { provideAnimations } from '@angular/platform-browser/animations';


// Function to customize the renderer for markdown
export function markedOptionsFactory(): MarkedOptions {
  const renderer = new MarkedRenderer();

  return {
    renderer: renderer,
    gfm: true,
    breaks: false,
    pedantic: false,
  };
}



export const appConfig: ApplicationConfig = {
  providers: [
    provideMarkdown({
      markedOptions: {
        provide: MARKED_OPTIONS,
        useFactory: markedOptionsFactory
      }
    }),
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideAnimations(),
    provideToastr({
      timeOut: 3000, // Toast will close after 3 seconds
      positionClass: 'toast-top-right', // Position of the toast
      preventDuplicates: true, // Prevent duplicate toasts
      progressBar: true, // Show progress bar
      closeButton: true, // Show close button
      tapToDismiss: false, // Disable tap to dismiss
    }), provideRouter(routes), provideHttpClient(withInterceptorsFromDi()),]
};
