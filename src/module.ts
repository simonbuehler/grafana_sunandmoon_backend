import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
import { ConfigEditor } from './components/ConfigEditor';
import { QueryEditor } from './components/QueryEditor';
import { SunAndMoonQuery, SunAndMoonDataSourceOptions } from './types';

export const plugin = new DataSourcePlugin<DataSource, SunAndMoonQuery, SunAndMoonDataSourceOptions>(DataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditor);
