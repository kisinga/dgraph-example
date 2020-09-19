import {
  trigger,
  state,
  style,
  transition,
  animate,
} from "@angular/animations";
import {
  Component,
  OnInit,
  ChangeDetectionStrategy,
  Input,
  Inject,
  AfterViewInit,
  OnChanges,
  ViewChild,
} from "@angular/core";
import {
  MatDialog,
  MatDialogRef,
  MAT_DIALOG_DATA,
} from "@angular/material/dialog";
import { MatPaginator } from "@angular/material/paginator";
import { MatSort } from "@angular/material/sort";
import { MatTableDataSource } from "@angular/material/table";
import { ActorsComponent } from "../actors/actors.component";
import { Actor, Film } from "../models/main";
export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
  description: string;
}

@Component({
  selector: "app-movies",
  templateUrl: "./movies.component.html",
  styleUrls: ["./movies.component.scss"],
  changeDetection: ChangeDetectionStrategy.OnPush,
  animations: [
    trigger("detailExpand", [
      state("collapsed", style({ height: "0px", minHeight: "0" })),
      state("expanded", style({ height: "*" })),
      transition(
        "expanded <=> collapsed",
        animate("225ms cubic-bezier(0.4, 0.0, 0.2, 1)")
      ),
    ]),
  ],
})
export class MoviesComponent implements AfterViewInit, OnChanges {
  @Input() films: Array<Film>;
  @Input() loading: boolean;
  @Input() phrase: string;
  dataSource: MatTableDataSource<Film>;

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  columnsToDisplay = ["uid", "name", "Directors", "length"];

  constructor(
    public dialogRef: MatDialogRef<MoviesComponent>,
    private dialog: MatDialog,
    @Inject(MAT_DIALOG_DATA) public data: Array<Film>
  ) {
    this.dataSource = new MatTableDataSource();

    this.dataSource.data = this.data;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }
  openDialog(film: Film): void {
    const dialogRef = this.dialog.open(ActorsComponent, {
      width: "75%",
      data: film.actors,
    });
  }
  ngOnChanges(changes: any) {
    this.dataSource.data = this.films;
  }
}
