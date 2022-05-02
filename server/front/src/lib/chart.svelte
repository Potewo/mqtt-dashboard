<script lang="ts">
  import Line from "svelte-chartjs/src/Line.svelte"
  let data = {
    labels: ["January", "February", "March", "April", "May", "June", "July"],
    datasets: [
      {
        label: "Sample",
        backgroundColor: "rgb(255, 99, 132)",
        borderColor: "rgb(255, 99, 132)",
        data: [0, 10, 5, 2, 20, 30, 45],
      },
    ],
  };
</script>
<div class="chart-container">
  <Line data={data} width={100} height={30}/>
</div>
<style>
  .chart-container {
    width: 80%;
    height: 40vh;
  }
</style>
