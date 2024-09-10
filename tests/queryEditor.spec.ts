import { test, expect } from '@grafana/plugin-e2e';

test.describe('QueryEditor', () => {
  test('should trigger a new query when Metric field is changed', async ({ panelEditPage, readProvisionedDataSource }) => {
    const ds = await readProvisionedDataSource({ fileName: 'datasources.yml' });
    await panelEditPage.datasource.set(ds.name);

    // Select metric
    await panelEditPage.getQueryEditorRow('A').getByLabel('Metric').selectOption({ label: 'Moon Altitude' });

    // Wait for a query request
    const queryReq = panelEditPage.waitForQueryDataRequest();
    await expect(await queryReq).toBeTruthy();
  });

  test('should trigger a new query when Latitude and Longitude fields are changed', async ({ panelEditPage, readProvisionedDataSource }) => {
    const ds = await readProvisionedDataSource({ fileName: 'datasources.yml' });
    await panelEditPage.datasource.set(ds.name);

    // Set Latitude
    await panelEditPage.getQueryEditorRow('A').getByRole('spinbutton', { name: 'Latitude' }).fill('45.0');
    // Set Longitude
    await panelEditPage.getQueryEditorRow('A').getByRole('spinbutton', { name: 'Longitude' }).fill('9.0');

    // Wait for a query request
    const queryReq = panelEditPage.waitForQueryDataRequest();
    await expect(await queryReq).toBeTruthy();
  });

  test('should return correct data in the Table visualization', async ({ panelEditPage, readProvisionedDataSource }) => {
    const ds = await readProvisionedDataSource({ fileName: 'datasources.yml' });
    await panelEditPage.datasource.set(ds.name);

    // Set Query Text or Metric
    await panelEditPage.getQueryEditorRow('A').getByRole('textbox', { name: 'Query Text' }).fill('moon_illumination');

    // Set Visualization to Table
    await panelEditPage.setVisualization('Table');

    // Refresh the panel
    await expect(panelEditPage.refreshPanel()).toBeOK();

    // Validate the response data (assuming moon_illumination returns specific values)
    await expect(panelEditPage.panel.data).toContainText(['moon']);
  });
});
