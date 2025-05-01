import { NgClass } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { AuthService } from '../../../core/services/auth.service';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-signup',
  imports: [ReactiveFormsModule, NgClass, RouterModule],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.scss'
})
export class SignupComponent {

  signUpForm: FormGroup;
  submitted = false;

  constructor(
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private router: Router,
    private toastr: ToastrService,
  ) {
    this.signUpForm = this.formBuilder.group({
      name: ['', [Validators.required, Validators.minLength(6)]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
    });
  }
  get f() {
    return this.signUpForm.controls;
  }
  onSubmit() {
    this.submitted = true;
    if (this.signUpForm.invalid) {
      return;
    }
    const payload = {
      name: this.signUpForm.value.name,
      email: this.signUpForm.value.email,
      password: this.signUpForm.value.password,
    };
    this.authService.SignUp(payload).subscribe({
      next: (response) => {
        console.log('SignUp successful:', response);
        this.toastr.success('SignUp successful!', 'Success');
        this.router.navigate(['/login']);
      },
      error: (error) => {
        console.log(error);
        this.toastr.error('SignUp failed. Please try again.', error.error.detail);
      },
    });
  }

}
