import { Component, OnInit, ChangeDetectionStrategy } from "@angular/core";
import { MoviesService } from "./movies.service";

@Component({
  selector: "app-movies",
  templateUrl: "./movies.component.html",
  styleUrls: ["./movies.component.scss"],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MoviesComponent implements OnInit {
  constructor(private moviesService: MoviesService) {}

  ngOnInit(): void {}
}
