import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { Actor, Film } from "./models/main";

@Injectable({
  providedIn: "root",
})
export class SearchService {
  constructor(private readonly httpClient: HttpClient) {}

  // searchMovies returns the movies matching the specified phrase
  // together with the related child components
  searchMovies(phrase: string): Observable<Film[]> {
    return this.generalSearch<Film>(phrase, "movies");
  }

  // searchActors returns the actors matching the specified phrase
  // together with the related child components
  searchActors(phrase: string): Observable<Actor[]> {
    return this.generalSearch<Actor>(phrase, "actors");
  }

  // Generic function the sends the query params to the same
  // rest endpoint
  private generalSearch<T>(
    phrase: string,
    searchtype: string
  ): Observable<T[]> {
    return this.httpClient.get<T[]>(`${environment.apiUrl}/api/search`, {
      params: { phrase, searchtype },
    });
  }
}
