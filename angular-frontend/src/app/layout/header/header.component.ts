import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../core/services/auth.service';
import { ChatEnums } from '../../core/enums/chat.enums';

@Component({
  selector: 'app-header',
  imports: [],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})
export class HeaderComponent {
  constructor(
    private router: Router,
    private authService: AuthService
  ) {

  }

  LogOut() {
    this.authService.LogOut();
    this.router.navigate(['/login'])
  }
  GoToDashBoard() {
    this.router.navigate(['/'])
  }
  GoToGeneralChat() {
    // this.router.navigate(['/general-chat'])
    this.router.navigate(['/chat-session'], {
      queryParams: {
        sessionId: ChatEnums.GENERAL_CHAT_SESSION,
      }
    });
  }
}
