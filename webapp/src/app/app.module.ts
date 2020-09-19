import { BrowserModule } from "@angular/platform-browser";
import { NgModule } from "@angular/core";

import { AppComponent } from "./app.component";
import { HttpClientModule } from "@angular/common/http";
import { MoviesComponent } from "./movies/movies.component";
import { ActorsComponent } from "./actors/actors.component";
import { MatSelectModule } from "@angular/material/select";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatInputModule } from "@angular/material/input";
import { MatCardModule } from "@angular/material/card";
import { MatButtonModule } from "@angular/material/button";
import { MatTableModule } from "@angular/material/table";
import { MatIconModule } from "@angular/material/icon";
import { MatProgressSpinnerModule } from "@angular/material/progress-spinner";
import { MatPaginatorModule } from "@angular/material/paginator";
import { MatSortModule } from "@angular/material/sort";
import {
  MatDialogModule,
  MatDialogRef,
  MAT_DIALOG_DATA,
} from "@angular/material/dialog";

@NgModule({
  declarations: [AppComponent, MoviesComponent, ActorsComponent],
  imports: [
    BrowserModule,
    HttpClientModule,
    MatSelectModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    BrowserAnimationsModule,
    MatCardModule,
    MatButtonModule,
    MatTableModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatPaginatorModule,
    MatSortModule,
    MatDialogModule,
  ],
  providers: [
    { provide: MAT_DIALOG_DATA, useValue: {} },
    { provide: MatDialogRef, useValue: {} },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
