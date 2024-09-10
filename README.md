# <img src="src/img/logo.svg" alt="logo" height="48px"/> Sun and Moon Datasource Backend Plugin for Grafana



## Overview

The **Sun and Moon Datasource Backend Plugin** for Grafana allows for metrics and events related to the sun and moon, such as solar noon, sunrise, sunset, moonrise, moonset, and more. This Go backend version is a port of the original [Grafana Sun and Moon Datasource](https://github.com/fetzerch/grafana-sunandmoon-datasource) using the [Suncalc Go library](https://github.com/sixdouglas/suncalc).

### Why This Plugin?

This backend version was developed to address the limitations of frontend-only plugins, particularly their incompatibility with **public dashboards** in Grafana. By handling the calculations on the backend, this plugin is fully functional for public dashboards, making it ideal for shared or embedded use cases.

## Features

- **Sun Events**: Solar noon, sunrise, sunset, golden hour, and other sun-related events.
- **Moon Events**: Moonrise, moonset, moon illumination, and more.
- **Backend Processing**: Moves the calculations to the backend, ensuring compatibility with public Grafana dashboards.

## Installation (while not available in the Grafana plugin repository)

1. **Download the Plugin**:
   - Download the latest version from [GitHub Releases](https://github.com/simonbuehler/sunandmoon_backend/releases).

2. **Install the Plugin**:
   - Place the plugin folder in the Grafana plugin directory:
     ```bash
     sudo cp -r simonbuehler-sunandmoon-datasource /var/lib/grafana/plugins/
     ```

3. **Configure Grafana**:
   - If required, allow loading unsigned plugins in your `grafana.ini`:
     ```ini
     [plugins]
     allow_loading_unsigned_plugins = simonbuehler-sunandmoon-datasource
     ```

4. **Restart Grafana**:
   - Restart the Grafana server to load the new plugin:
     ```bash
     sudo systemctl restart grafana-server
     ```

## Usage

1. **Add the Datasource**:
   - Go to **Configuration > Data Sources** in Grafana, click **Add Data Source**, and select **Sun and Moon (backend version)** from the list.

2. **Configure the Datasource**:
   - Set default latitude and longitude for sun and moon calculations. These can also be overridden on a per-query basis.

3. **Create Panels**:
   - Add panels and choose sun and moon metrics like moon illumination or solar noon to visualize your data.

## Credits

- **Original Plugin**: [fetzerch/grafana-sunandmoon-datasource](https://github.com/fetzerch/grafana-sunandmoon-datasource)
- **Suncalc Library**: [sixdouglas/suncalc](https://github.com/sixdouglas/suncalc)

## License

This project is licensed under the [MIT License](LICENSE).
