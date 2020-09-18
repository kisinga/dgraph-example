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
} from "@angular/core";
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
export class MoviesComponent implements OnInit {
  @Input() movies: PeriodicElement;
  @Input() loading: boolean;
  @Input() phrase: string;

  columnsToDisplay = ["name", "weight", "symbol", "position"];
  expandedElement: PeriodicElement | null;
  constructor() {}

  ngOnInit(): void {}
}
