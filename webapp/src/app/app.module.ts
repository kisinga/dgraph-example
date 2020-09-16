import { BrowserModule } from "@angular/platform-browser";
import { NgModule } from "@angular/core";

import { AppComponent } from "./app.component";
import { TechComponent } from "./tech/tech.component";
import { HttpClientModule } from "@angular/common/http";
import { MoviesComponent } from "./movies/movies.component";
import { ActorsComponent } from "./actors/actors.component";
import { MatSlideToggleModule } from "@angular/material/slide-toggle";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

@NgModule({
  declarations: [AppComponent, TechComponent, MoviesComponent, ActorsComponent],
  imports: [BrowserModule, HttpClientModule, MatSlideToggleModule, BrowserAnimationsModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
