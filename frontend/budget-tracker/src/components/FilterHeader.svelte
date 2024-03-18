<script>
    import { selectionMode, selectedStartDate, selectedEndDate, selectedMonth, selectedYear, selectedDay } from "../stores/selectionMode";
    import FaRegArrowAltCircleLeft from 'svelte-icons/fa/FaRegArrowAltCircleLeft.svelte'
    import FaRegArrowAltCircleRight from 'svelte-icons/fa/FaRegArrowAltCircleRight.svelte'
    import FaArrowRight from 'svelte-icons/fa/FaArrowRight.svelte'

    const months = [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December"
    ];

    function formatTime(date) {
        return date.toISOString().split('T')[0];
    }

    function formatMonth(number) {
        return months[number];
    }

    function clearSelection() {
        selectionMode.set(null);
        selectedStartDate.set(null);
        selectedEndDate.set(null);
        selectedMonth.set(null);
        selectedYear.set(null);
        selectedDay.set(null);
    }

    function nextYear() {
        selectedYear.update(n => n + 1);
    }

    function previousYear() {
        selectedYear.update(n => n - 1);
    }

    function nextMonth() {
        selectedMonth.update(n => {
            if (n >= 11) {
                selectedYear.update(y => y + 1);
                return 0;
            } else {
                return n + 1;
            }
        });
    }

    function previousMonth() {
    selectedMonth.update(n => {
        if (n <= 0) {
            selectedYear.update(y => y - 1);
            return 11;
        } else {
            return n - 1;
        }
    });
}
</script>

{#if $selectionMode === null}
    <div class="header-null"></div>
{:else if $selectionMode === 'range'}
    <div class="header range">
        {#if $selectedStartDate && $selectedEndDate}
        <div class="header-container">
            <div>{formatTime($selectedStartDate)}</div>
            <span class="icon"><FaArrowRight/></span>
            <div>{formatTime($selectedEndDate)}</div>
            <button class="icon-clear" on:click={clearSelection}>Clear</button>
        </div>
        {:else}
            <div>No range selected</div>
        {/if}
    </div>
{:else if $selectionMode === 'year'}
    <div class="header year">
        <div class="header-container">
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <span on:click={previousYear} class="icon-nav"><FaRegArrowAltCircleLeft/></span>
            <div class="year alone">{$selectedYear}</div>
             <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <span on:click={nextYear} class="icon-nav"><FaRegArrowAltCircleRight/></span>
            <button class="icon-clear" on:click={clearSelection}>Clear</button>
        </div>

    </div>
    
{:else if $selectionMode === 'month'}
    <div class="header month">
        <div class="header-container">
             <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <span on:click={previousMonth} class="icon-nav"><FaRegArrowAltCircleLeft/></span>
            <span class="month">{formatMonth($selectedMonth)}</span>
            <span class="year">{$selectedYear}</span>
             <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <span on:click={nextMonth} class="icon-nav"><FaRegArrowAltCircleRight/></span>
            <button class="icon-clear" on:click={clearSelection}>Clear</button>
        </div>
    </div>
{:else if $selectionMode === 'day'}
    <div class="header">Selected day: {$selectedDay}</div>
    <button class="icon-clear" on:click={clearSelection}>Clear</button>
{/if}

<style>
     .icon {
        width: 24px;
        height: 24px;
        margin: 20px;
        padding-bottom: 20px;
        margin-top:42px;
    }

    .icon-nav {
        width: 24px;
        height: 24px;
        margin: 20px;
        margin-top: 40px;
        padding-bottom: 20px;
        cursor: pointer;
    }

    .icon-nav:hover {
        color: #3A87F2;
    }

    .icon-clear {

        margin: 20px;
        cursor: pointer;
    }

    .month {
        font-size: 1.5em;
        margin-right: 20px;
    }

    .year {
        font-size: 1.5em;
        margin-right: 20px;
    }

    .year.alone {
        font-size: 1.5em;
    }
    .header {
        display: flex;
        font-size: 1.5em;
        margin-bottom: 20px;
        color: black;
        background-color: #E8EFF1;
        justify-content: center;
        align-items: center;
        padding: 20px;
        position: fixed;
        width: calc(100% - 300px);
        top: 0;
        z-index: 1000;
        height: 40px;
        justify-content: center;
    }

    .header-container {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .header-null {
        display: none;
    }
</style>