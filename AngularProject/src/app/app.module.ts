import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '@angular/material';
import 'hammerjs';

import { AvailableRoutes } from './app.routing';

import { AppComponent } from './app.component';
import { TimerComponent } from './timer/timer.component';
import { CreateComponent } from './create/create.component';
import { LoadComponent } from './load/load.component';

@NgModule({
    declarations: [
        AppComponent,
        TimerComponent,
        CreateComponent,
        LoadComponent
    ],
    imports: [
        BrowserModule,
        FormsModule,
        HttpModule,
        RouterModule,
        RouterModule.forRoot(AvailableRoutes),
        MaterialModule.forRoot()
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule { }
