<script>
    import Calendar from "../../components/Calendar.svelte";
    import IncomeTable from "../../components/IncomeTable.svelte";
    import AddIncome from "../../components/AddIncome.svelte";
    import MonthlyIncomeGraph from "../../components/MonthlyIncomeGraph.svelte";
</script>

<style>

    .dashboard-container {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-template-rows: auto auto;
        gap: 20px;
        padding: 40px;
        margin: 0 auto;
        max-width: 1400px;
    }


    .add-income-container,
    .calendar-container,
    .bar-chart-container,
    .budget-table-container {
        background: #FFFFFF;
        padding: 20px;
        background-color: transparent;
        
    }

    .bar-chart-container {
        grid-column: 1 / 3;
        grid-row: 1;
    }

    .calendar-container {
        grid-column: 3 / 4;
        grid-row: 1;
    }


    .add-income-container {
        grid-column: 3 / 4;
        grid-row: 2;
        text-align: right;
        align-items: center;
    }


    .budget-table-container {
        grid-column: 1 / 3;
        grid-row: 2 / 4;
        padding-top: 0%;
    }

   

    @media (max-width: 1400px) {
        .dashboard-container {
            grid-template-columns: 1fr;
            grid-template-rows: auto auto auto auto;
        }

        .calendar-container {
            grid-column: 1 / 2;
            grid-row: 1;
        }

        .bar-chart-container {
            grid-column: 1 / 2;
            grid-row: 2;
        }

        .add-income-container {
            grid-column: 1 / 2;
            grid-row: 3;
            text-align: center;
        }

        .budget-table-container {
            grid-column: 1 / 2;
            grid-row: 4;
        }
    }
</style>


<div class="dashboard-container">
    <div class="bar-chart-container">
        <MonthlyIncomeGraph />
    </div>
   
    <div class="add-income-container">
        <AddIncome />
    </div>
 
    <div class="budget-table-container">
        <IncomeTable />
    </div>
    <div class="calendar-container">
        <Calendar />
    </div>

</div>
