import { DataSourceJsonData } from '@grafana/data';
import { DataQuery } from '@grafana/schema';

// Typ für die Abfragen, die an das Backend gesendet werden
export interface SunAndMoonQuery extends DataQuery {
  target?: string[]; // Array von Metriken, die abgefragt werden
  latitude?: string; // Optional: Breitenangabe als String (für Eingaben im Editor)
  longitude?: string; // Optional: Längenangabe als String (für Eingaben im Editor)
}

// Standardwerte für Abfragen (Metriken und ggf. Default-Latitude/Longitude)
export const DEFAULT_QUERY: Partial<SunAndMoonQuery> = {
  target: ['moon_illumination'], // Standard-Metrik für die Abfrage
};

// Typ für einen einzelnen Datenpunkt (Zeit und Wert)
export interface DataPoint {
  time: number; // Zeitstempel als Zahl (Unix-Zeit)
  value: number; // Wert des Datenpunkts
}

// Antwortstruktur, die vom Backend zurückgegeben wird
export interface DataSourceResponse {
  datapoints: DataPoint[]; // Array von Datenpunkten
}

/**
 * Optionen, die für jede Datenquelleninstanz konfiguriert werden.
 * Diese Optionen können in Grafana für jede Datenquelle individuell eingestellt werden.
 */
export interface SunAndMoonDataSourceOptions extends DataSourceJsonData {
  latitude?: number; // Optional: Breitenangabe (Wird als Zahl gespeichert)
  longitude?: number; // Optional: Längenangabe (Wird als Zahl gespeichert)
}
