<script lang="ts">
  export let selected: boolean = false;
  export let type: string;
  export let value: number;
  const typeDisplayNameTable = {
    "temperature": "温度",
    "humidity": "湿度",
    "airPressure": "気圧",
  }
  const typeUnitTable = {
    "temperature": "℃",
    "humidity": "%",
    "airPressure": "hPa",
  }
</script>
<div class="my-card my-card" class:my-card-selected={selected}>
  <p class="my-card-title">{typeDisplayNameTable[type]}</p>
  <p class="my-card-value">{value}{typeUnitTable[type]}</p>
</div>

<style>
  .my-card {
    width: 30%;
    border-radius: 10px;
    max-width: 200px;
    padding: 30px;
    text-align: center;
    margin: 20px;
    box-shadow: 0 2px 5px #ccc;
  }

  .my-card-value {
    font-size: 3rem;
  }

  .my-card-selected {
    border: solid 2px #ccf;
  }
</style>
