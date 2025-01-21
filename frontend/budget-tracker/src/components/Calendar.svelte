<script>
  import FaAngleLeft from 'svelte-icons/fa/FaAngleLeft.svelte'
  import FaAngleRight from 'svelte-icons/fa/FaAngleRight.svelte'
  import { onMount } from 'svelte';
  import { get } from 'svelte/store';
  import { selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate, selectedDay } from '../stores/selectionMode';
  import { derived } from 'svelte/store';
  import { fromISOString } from '../utility/functions';

  onMount(() => {
    updateCalendar(); // If this prepares the UI
  });


  let currentYear = new Date().getFullYear();
  let currentMonth = new Date().getMonth(); 
  let selectedDate = new Date();
  let yearViewActive = false;
  let rangeViewActive = false;
  let yearsToShow = [];


  let rangeStart = null;
  let rangeEnd = null;

  $: $selectionMode, updateCalendar();

  $: console.log($selectionMode)

  const firstDay = new Date(currentYear, currentMonth, 1).getDay();

  const daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate();

  let dates = [];
  for (let i = 0; i < firstDay; i++) {
    dates.push(null);
  }
  for (let day = 1; day <= daysInMonth; day++) {
    dates.push(new Date(currentYear, currentMonth, day));
  }

  // Generate array of years (e.g., current year ± 6 years)
  for (let i = -6; i <= 6; i++) {
      yearsToShow.push(currentYear + i);
  }

  function clearSelection() {
    selectedYear.set(currentYear);
    selectedMonth.set(currentMonth);
    selectedDay.set(null);
    selectedStartDate.set(null);
    selectedEndDate.set(null);
    selectionMode.set("month");
    selectToday();
    updateCalendar();
    //FetchExpenses();
    //fetchIncome();
    yearViewActive = false;
    rangeViewActive = false;
  }

  function toggleSelectionMode(mode) {
        selectionMode.update(current => {
            if (current === mode) {
                clearSelection();
                yearViewActive = false;
                rangeViewActive = false;
                return null;
            } else {
                switch (mode) {
                    case 'day':
                        selectedDay.set(selectDate(selectedDate));
                        yearViewActive = false;
                        rangeViewActive = false;
                        break;
                    case 'month':
                        selectedMonth.set(currentMonth);
                        selectedYear.set(currentYear);
                        yearViewActive = false;
                        rangeViewActive = false;
                        break;
                    case 'year':
                        selectedYear.set(currentYear);
                        yearViewActive = true;
                        rangeViewActive = false;
                        break;
                    case 'range':
                        selectedStartDate.set(rangeStart);
                        selectedEndDate.set(rangeEnd);
                        yearViewActive = false;
                        rangeViewActive = true;
                        break;
                }
                return mode;
            }
        });
    }


  function jumpToMonth(monthIndex) {
        currentMonth = monthIndex;
        if (get(selectionMode) === 'month') {
            selectedMonth.set(monthIndex);
        }
        updateCalendar();
    }

    function jumpYears(direction) {
        currentYear += (direction * 12);
        if (get(selectionMode) === 'year') {
            selectedYear.set(currentYear);
        }
        updateCalendar();
    }

  function selectToday() {
      const today = new Date();
      selectedDate = today;

      // Check if the year or month has changed
      if (today.getFullYear() !== currentYear || today.getMonth() !== currentMonth) {
          currentYear = today.getFullYear();
          currentMonth = today.getMonth();
          updateCalendar(); 
      }
  }

  function selectDate(date) {
  const currentMode = get(selectionMode);
  if (currentMode === 'day') {
    selectedDay = date;
  }
  if (currentMode === 'range') {
      let start = get(selectedStartDate);
      let end = get(selectedEndDate);

      if (start === null || (start !== null && end !== null)) {
          selectedStartDate.set(date);
          selectedEndDate.set(null);
      } else if (start !== null && end === null) {
      if (date < start) {
          selectedStartDate.set(start);
          selectedEndDate.set(date);
      } else {
          selectedEndDate.set(date);
        }
      }
      } else {
      if (date && date.getMonth() < currentMonth || (date.getMonth() === 11 && currentMonth === 0 && date.getFullYear() < currentYear)) {
          currentMonth--;
          if (currentMonth < 0) {
          currentMonth = 11;
          currentYear--;
          }
      }
      else if (date && date.getMonth() > currentMonth || (date.getMonth() === 0 && currentMonth === 11 && date.getFullYear() > currentYear)) {
          currentMonth++;
          if (currentMonth > 11) {
          currentMonth = 0;
          currentYear++;
          }
      }
    }

  selectedDate = date;
  updateCalendar(); 
  }

  function selectYear(year) {
        currentYear = year;
        selectedYear.set(year);
        updateCalendar();
    }

  function updateCalendar() {
  const firstDayIndex = new Date(currentYear, currentMonth, 1).getDay();
  const prevMonthDays = new Date(currentYear, currentMonth, 0).getDate();
  const daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate();

  dates = [];

  // Calculate leading days for the previous month to fill the first week
  let tempDates = [];
  for (let i = 0; i < firstDayIndex; i++) {
      const day = prevMonthDays - (firstDayIndex - i - 1);
      tempDates.push(new Date(currentYear, currentMonth - 1, day));
  }

  tempDates = tempDates.reverse();

  tempDates.forEach(date => {
  dates.unshift(date);
  });


  // Current month's days
  for (let day = 1; day <= daysInMonth; day++) {
      dates.push(new Date(currentYear, currentMonth, day));
  }

  // Calculate leading days for the next month to fill the last week
  let daysToAddForNextMonth = 7 - (dates.length % 7);
  if (daysToAddForNextMonth && daysToAddForNextMonth < 7) {
      for (let i = 1; i <= daysToAddForNextMonth; i++) {
          dates.push(new Date(currentYear, currentMonth + 1, i));
      }
  }
}

  function previousMonth() {
    const currentMode = get(selectionMode);
    currentMonth--;
    if (currentMonth < 0) {
      currentMonth = 11;
      currentYear--;
    }
    if (currentMode === 'month') {
      selectedMonth.set(currentMonth);
      selectedYear.set(currentYear);
    }
    updateCalendar();
  }

  function nextMonth() {
    const currentMode = get(selectionMode);
    currentMonth++;
    if (currentMonth > 11) {
      currentMonth = 0;
      currentYear++;
    }
    if (currentMode === 'month') {
      selectedMonth.set(currentMonth);
      selectedYear.set(currentYear);
    }
    updateCalendar();
  }

  function previousYear() {
    const currentMode = get(selectionMode);
    currentYear--;
    if (currentMode === 'year') {
      selectedYear.set(currentYear);
    }
    updateCalendar();
  }

  function nextYear() {
    const currentMode = get(selectionMode);
    currentYear++;
    if (currentMode === 'year') {
      selectedYear.set(currentYear);
    }
    updateCalendar();
  }

  const modeLabel = derived(selectionMode, ($selectionMode) => {
  switch ($selectionMode) {
      case 'year':
        return 'Year Mode';
      case 'range':
        return 'Custom Range';
      default:
        return 'Month Mode';
    }
  });
</script>

<div class="calendar-box">
  <div class="selector">
    <select 
        class="select-dropdown"
        on:change={(e) => toggleSelectionMode(e.target.value)}
        bind:value={$selectionMode}>
        <option value="month">Monthly View</option>
        <option value="year">Yearly View</option>
        <option value="range">Custom Range</option>
    </select>
    <div class="mode-badge">{$modeLabel}</div>
  </div>


  {#if yearViewActive}
        <!-- Year View -->
        <div class="year-view">
            <div class="year-navigation">
                <button class="icon-button" on:click={() => jumpYears(-1)}>
                    <span class="prev-year-arrow"><FaAngleLeft/></span>
                </button>
                <span class="year-range">{yearsToShow[0]} - {yearsToShow[yearsToShow.length - 1]}</span>
                <button class="icon-button" on:click={() => jumpYears(1)}>
                    <span class="next-year-arrow"><FaAngleRight/></span>
                </button>
            </div>
            <div class="years-grid">
                {#each yearsToShow as year}
                    <button 
                        class="year-cell {year === $selectedYear ? 'selected' : ''}"
                        on:click={() => selectYear(year)}>
                        {year}
                    </button>
                {/each}
            </div>
        </div>
    {:else}
    {#if rangeViewActive}
    <div class="date-range">
        {#if $selectedStartDate && $selectedEndDate}
            {fromISOString($selectedStartDate)} 
            - 
            {fromISOString($selectedEndDate)}
        {:else if $selectedStartDate}
            {fromISOString($selectedStartDate)}
        {:else}
            Select date range
        {/if}
    </div>
  {:else}
  <div class="month-view">
    <div class="month-navigation">
      <button class="icon-button" on:click={previousMonth}>
        <span class="prev-month-arrow"><FaAngleLeft/></span>
      </button>
      <div class="month-button
      ">{
          (new Date(currentYear, currentMonth).toLocaleString('default', { month: 'long', year: 'numeric' })
            .replace(/^\w/, (c) => c.toUpperCase()))
        }</div>
    
      <button class="icon-button" on:click={nextMonth}>
        <span class="next-month-arrow"><FaAngleRight/></span>
      </button>
    </div>
</div>
{/if}
<div class="the-calendar">
<div class="calendar-days">
    <p>Sun</p><p>Mon</p><p>Tue</p><p>Wed</p><p>Thu</p><p>Fri</p><p>Sat</p>
</div>
  <div class="calendar-container">
    {#each dates as date, index (date ? date.toDateString() : 'empty-' + index)}
    <button 
        class="day {date?.toDateString() === selectedDate.toDateString() ? 'selected' : ''}
                  {get(selectedStartDate) && get(selectedEndDate) && date >= get(selectedStartDate) && date <= get(selectedEndDate) ? 'in-range' : ''}
                  {date?.getMonth() !== currentMonth ? 'previous-days' : ''}" 
        on:click={() => selectDate(date)}>
        {date ? date.getDate() : ''}
    </button>
{/each}
  </div>
</div>

{/if}
  <div class="dates-below">
  <button class="clear-button" on:click={() => clearSelection()}>Clear</button>
  </div>  
</div>


<style>
.selector {
  display: flex;
  align-items: center;
  gap: 12px;
}

.date-range {
    text-align: center;
    padding: 10px;
    font-size: 1rem;
    color: var(--text-color);
    font-weight: 500;
}

/* Select dropdown styling */
.select-dropdown {
  width: 85%;
  padding: 8px 12px;
  font-size: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #fff;
  color: #333;
  appearance: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
  cursor: pointer;
}

.select-dropdown:focus {
  border-color: #4a86ff;
  box-shadow: 0 0 4px rgba(74, 134, 255, 0.5);
  outline: none;
}

.mode-badge {
  padding: 6px 12px;
  font-size: 0.875rem;
  font-weight: 500;
  border-radius: 12px;
  background-color: #f1f1f1;
  color: #666;
  white-space: nowrap;
  text-transform: capitalize;
}

.select-dropdown:hover {
  border-color: #4a86ff;
}


  .calendar-box {
    width: 100%;
    max-width: 20rem;
    margin-top: 40px;
    padding: 20px;

    background-color: var(--sidenav-color);
    font-family: var(--font-family);
  }

  .calendar-container {
    display: grid;
    grid-template-columns: repeat(7, 1fr);
    grid-gap: 4px;
    padding-top: 5px;
    background-color: var(--sidenav-color);


    
  }

  .calendar-days {
      display: grid;
      grid-template-columns: repeat(7, 1fr);
      text-align: center;
      font-size: 12px; 
      color: var(--text-color);
      font-weight: bold;
      margin-left: 0;
      margin-right: 0;
      padding-left: 0;
      padding-right: 0;
  }


  .month-button {
      background-color: transparent;
      border: none;
      color: black;
      cursor: pointer; 
      font-size: 1rem; 
      transition: transform 0.1s ease; 
      }

  .month-button:hover {
      transform: bold;
  }

  .previous-days {
      color: #ccc;
  }
  
  .day {
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 10px;
      background: transparent;
      border-radius: 10px;
      cursor: pointer;
      transition: background-color 0.1s ease;
      border: 0;
  }

  .day.in-range {
      background-color: var(--primary-button-color);
      color: white;
  }
  
  .day:hover {
      background-color: #68a1f1;
  }

  .selected {
      background-color: #3A87F2;
      color: white;
  }

  .icon-button {
      background: none;
      border: none;
      padding: 0;
      cursor: pointer;
      outline: inherit;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      color: white;
  }

  .prev-month-arrow, .next-month-arrow, .prev-year-arrow, .next-year-arrow {
      width: 20px;
      height: 20px;
      color: black;
  }

  .prev-month-arrow:hover, .next-month-arrow:hover, .prev-year-arrow:hover, .next-year-arrow:hover {
      color: #68a1f1;
  }

  .dates-below {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-top: 10px;
      margin-left: 15px;
      margin-right: 15px;
  }
  .month-view,
  .year-view {
        padding: 1rem 0;
    }
    .month-navigation,
    .year-navigation {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 1rem;
        padding: 0 1rem;
    }

    .year-range {
        font-weight: 500;
        color: var(--text-color);
    }

    .years-grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: 0.5rem;
        padding: 0.5rem;
    }

    .year-cell {
        padding: 0.75rem;
        text-align: center;
        background: transparent;
        border: none;
        border-radius: 8px;
        cursor: pointer;
        transition: all 0.2s ease;
    }

    .year-cell:hover {
        background-color: var(--primary-button-color);
        color: white;
    }

    .year-cell.selected {
        background-color: var(--primary-button-color);
        color: white;
    }



  .clear-button {
      background-color: var(--primary-button-color);
      color: white;
      padding: 5px;
      border-radius: 5px;
      cursor: pointer;
      border: 0;
  }

</style>