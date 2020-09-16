import { Component } from "@angular/core";
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from "@angular/forms";
import { PeriodicElement } from "./movies/movies.component";
import { delay, takeUntil } from "rxjs/operators";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.scss"],
})
export class AppComponent {
  title = "DGraph-Example";

  searchType: string[] = ["Movie Name", "Actor"];
  searchControl = new FormControl(this.searchType[0], [Validators.required]);
  phraseControl = new FormControl("", [Validators.required]);
  selectedSearchType = this.searchType[0];
  searchForm: FormGroup;
  constructor(fb: FormBuilder) {
    this.searchForm = fb.group({
      search: this.searchControl,
      phrase: this.phraseControl,
    });
    this.searchControl.valueChanges.subscribe((c) => {
      this.selectedSearchType = c;
    });
    this.phraseControl.valueChanges.pipe(delay(50)).subscribe((c) => {});
  }

  ELEMENT_DATA: PeriodicElement[] = [
    {
      position: 1,
      name: "Hydrogen",
      weight: 1.0079,
      symbol: "H",
      description: `Hydrogen is a chemical element with symbol H and atomic number 1. With a standard
        atomic weight of 1.008, hydrogen is the lightest element on the periodic table.`,
    },
    {
      position: 2,
      name: "Helium",
      weight: 4.0026,
      symbol: "He",
      description: `Helium is a chemical element with symbol He and atomic number 2. It is a
        colorless, odorless, tasteless, non-toxic, inert, monatomic gas, the first in the noble gas
        group in the periodic table. Its boiling point is the lowest among all the elements.`,
    },
    {
      position: 3,
      name: "Lithium",
      weight: 6.941,
      symbol: "Li",
      description: `Lithium is a chemical element with symbol Li and atomic number 3. It is a soft,
        silvery-white alkali metal. Under standard conditions, it is the lightest metal and the
        lightest solid element.`,
    },
    {
      position: 4,
      name: "Beryllium",
      weight: 9.0122,
      symbol: "Be",
      description: `Beryllium is a chemical element with symbol Be and atomic number 4. It is a
        relatively rare element in the universe, usually occurring as a product of the spallation of
        larger atomic nuclei that have collided with cosmic rays.`,
    },
    {
      position: 5,
      name: "Boron",
      weight: 10.811,
      symbol: "B",
      description: `Boron is a chemical element with symbol B and atomic number 5. Produced entirely
        by cosmic ray spallation and supernovae and not by stellar nucleosynthesis, it is a
        low-abundance element in the Solar system and in the Earth's crust.`,
    },
  ];
}
