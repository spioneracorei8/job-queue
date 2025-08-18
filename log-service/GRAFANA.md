# Step-by-Step: Connect Grafana with Loki

## 1. Open Grafana
- Navigate to your Grafana instance in your browser.

## 2. Add Loki Data Source
- Go to **Menu** > **Connections** > **Data sources** > **Add data source**.

## 3. Select Loki
- In the list of data sources, select **Loki**.

## 4. Configure Loki URL
- Set the **URL** field to:  
    ```
    http://loki:3100
    ```

## 5. Save & Test
- Click **Save & Test** to verify the connection.

---

Your Grafana instance is now connected to Loki!