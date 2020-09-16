import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { Movie } from "./models/movie";

@Injectable({
  providedIn: "root",
})
export class MainService {
  constructor(private readonly httpClient: HttpClient) {}

  searchMovies(): Observable<Movie[]> {
    return this.httpClient.get<Movie[]>(`${environment.apiUrl}/api/movies`);
  }
  searchActors(): Observable<Movie[]> {
    return this.httpClient.get<Movie[]>(`${environment.apiUrl}/api/actors`);
  }
}
