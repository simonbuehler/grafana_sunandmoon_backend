import { test, expect } from '@grafana/plugin-e2e';
import { SunAndMoonDataSourceOptions } from '../src/types';

test('"Save & test" should be successful when latitude and longitude are valid', async ({
  createDataSourceConfigPage,
  readProvisionedDataSource,
  page,
}) => {
  const ds = await readProvisionedDataSource<SunAndMoonDataSourceOptions>({ fileName: 'datasources.yml' });
  const configPage = await createDataSourceConfigPage({ type: ds.type });

  // Set valid latitude and longitude values
  await page.getByLabel('Latitude').fill(ds.jsonData.latitude?.toString() ?? '48.3984');
  await page.getByLabel('Longitude').fill(ds.jsonData.longitude?.toString() ?? '9.9910');
  
  // Save and test should be OK
  await expect(configPage.saveAndTest()).toBeOK();
});

test('"Save & test" should fail when latitude is missing', async ({
  createDataSourceConfigPage,
  readProvisionedDataSource,
  page,
}) => {
  const ds = await readProvisionedDataSource<SunAndMoonDataSourceOptions>({ fileName: 'datasources.yml' });
  const configPage = await createDataSourceConfigPage({ type: ds.type });

  // Leave the latitude empty and set only the longitude
  await page.getByLabel('Latitude').fill('');
  await page.getByLabel('Longitude').fill(ds.jsonData.longitude?.toString() ?? '9.9910');
  
  // Save and test should not be OK
  await expect(configPage.saveAndTest()).not.toBeOK();
});

test('"Save & test" should fail when longitude is missing', async ({
  createDataSourceConfigPage,
  readProvisionedDataSource,
  page,
}) => {
  const ds = await readProvisionedDataSource<SunAndMoonDataSourceOptions>({ fileName: 'datasources.yml' });
  const configPage = await createDataSourceConfigPage({ type: ds.type });

  // Set only latitude and leave longitude empty
  await page.getByLabel('Latitude').fill(ds.jsonData.latitude?.toString() ?? '48.3984');
  await page.getByLabel('Longitude').fill('');
  
  // Save and test should not be OK
  await expect(configPage.saveAndTest()).not.toBeOK();
});
