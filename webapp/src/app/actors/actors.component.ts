import { OnChanges } from "@angular/core";
import { AfterViewInit } from "@angular/core";
import {
  Component,
  ChangeDetectionStrategy,
  Input,
  ViewChild,
} from "@angular/core";
import { MatPaginator } from "@angular/material/paginator";
import { MatSort } from "@angular/material/sort";
import { MatTableDataSource } from "@angular/material/table";
import { Actor } from "../models/main";

@Component({
  selector: "app-actors",
  templateUrl: "./actors.component.html",
  styleUrls: ["./actors.component.scss"],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ActorsComponent implements AfterViewInit, OnChanges {
  @Input() actors: Array<Actor>;
  @Input() phrase: string;
  @Input() loading: boolean;
  dataSource: MatTableDataSource<Actor>;

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  columnsToDisplay = ["uid", "name", "length"];
  constructor() {
    this.dataSource = new MatTableDataSource();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  ngOnChanges(changes: any) {
    this.dataSource.data = this.actors;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }
}
