// This has simple helpers to help with type safety
export enum SearchType {
  "Movie Name",
  "Actor",
}

export const SearchIds = Object.keys(SearchType)
  .filter((key) => isNaN(Number(SearchType[key])))
  .map((t) => Number(t));

export const SearchMap = Object.keys(SearchType)
  .filter((key) => isNaN(Number(SearchType[key])))
  .map((t) => {
    return {
      pos: t,
      value: SearchType[t],
    };
  });
export const SearchTypeNames = Object.keys(SearchType).filter(
  (key) => !isNaN(Number(SearchType[key]))
);
