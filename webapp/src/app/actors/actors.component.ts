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
import { PeriodicElement } from "../movies/movies.component";

@Component({
  selector: "app-actors",
  templateUrl: "./actors.component.html",
  styleUrls: ["./actors.component.scss"],
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
export class ActorsComponent implements OnInit {
  @Input() movies: PeriodicElement;
  @Input() phrase: string;
  columnsToDisplay = ["name", "weight", "symbol", "position"];
  expandedElement: PeriodicElement | null;
  constructor() {}

  ngOnInit(): void {}
}
