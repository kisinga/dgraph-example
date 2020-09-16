import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'app-actors',
  templateUrl: './actors.component.html',
  styleUrls: ['./actors.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ActorsComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
