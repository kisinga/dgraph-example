import { Inject, OnChanges } from "@angular/core";
import { AfterViewInit } from "@angular/core";
import {
  Component,
  ChangeDetectionStrategy,
  Input,
  ViewChild,
} from "@angular/core";
import { MatDialog, MAT_DIALOG_DATA } from "@angular/material/dialog";
import { MatPaginator } from "@angular/material/paginator";
import { MatSort } from "@angular/material/sort";
import { MatTableDataSource } from "@angular/material/table";
import { Actor, Film } from "../models/main";
import { MoviesComponent } from "../movies/movies.component";

@Component({
  selector: "app-actors",
  templateUrl: "./actors.component.html",
  styleUrls: ["./actors.component.scss"],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ActorsComponent implements AfterViewInit, OnChanges {
  @Input() actors: Array<Actor>;
  @Input() loading: boolean;
  @Input() root: boolean;

  dataSource: MatTableDataSource<Actor>;

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  columnsToDisplay = ["uid", "name"];
  constructor(
    private dialog: MatDialog,
    @Inject(MAT_DIALOG_DATA) public data: Array<Actor>
  ) {
    this.dataSource = new MatTableDataSource();
    this.dataSource.data = data;
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
    if (this.root) {
    }
  }

  ngOnChanges(changes: any) {
    if (this.root) {
      if (this.columnsToDisplay.length === 2) {
        this.columnsToDisplay[2] = "length";
      }
      this.dataSource.data = this.actors;
    }
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }
  openDialog(actor: Actor): void {
    const dialogRef = this.dialog.open(MoviesComponent, {
      width: "75%",
      data: actor.films_acted,
    });
  }
}
