<script>
    import FaAngleLeft from 'svelte-icons/fa/FaAngleLeft.svelte'
    import FaAngleRight from 'svelte-icons/fa/FaAngleRight.svelte'
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';
    import { selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate, selectedDay } from '../stores/selectionMode';
    

    function clearSelection() {
      selectedYear.set(null);
      selectedMonth.set(null);
      selectedDay.set(null);
      selectedStartDate.set(null);
      selectedEndDate.set(null);
      selectionMode.set(null);
      selectToday();
      updateCalendar();
    }

    function toggleSelectionMode(mode) {
    selectionMode.update(current => {
        if (current === mode) {
            clearSelection();
        } else {
            switch (mode) {
                case 'day':
                    selectedDay.set(selectDate(selectedDate));
                    break;
                case 'month':
                    selectedMonth.set(currentMonth);
                    selectedYear.set(currentYear);
                    break;
                case 'year':
                    selectedYear.set(currentYear);
                    break;
                case 'range':
                    selectedStartDate.set(rangeStart);
                    selectedEndDate.set(rangeEnd);
                    break;
            }
            return mode;
        }
    });
}

    onMount(() => {
        updateCalendar();
    });

    let currentYear = new Date().getFullYear();
    let currentMonth = new Date().getMonth(); 
    let selectedDate = new Date();


    let rangeStart = null;
    let rangeEnd = null;

    $: $selectionMode, updateCalendar();

    const firstDay = new Date(currentYear, currentMonth, 1).getDay();
  
    const daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate();
  
    let dates = [];
    for (let i = 0; i < firstDay; i++) {
      dates.push(null);
    }
    for (let day = 1; day <= daysInMonth; day++) {
      dates.push(new Date(currentYear, currentMonth, day));
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

  </script>
  
  <div class="calendar-box">
    <div class="calendar-controls">
      <div class="calendar-months">
        <button class="icon-button" on:click={previousMonth}>
          <span class="sr-only">Left Arrow</span>
            <span class="prev-month-arrow"><FaAngleLeft/></span>
        </button>
        <button class="month-button {$selectionMode === 'month' ? 'active' : ''}
        " on:click={() => toggleSelectionMode('month')}>{
            (new Date(currentYear, currentMonth).toLocaleString('default', { month: 'long' })
              .replace(/^\w/, (c) => c.toUpperCase()))
          }</button>
      
        <button class="icon-button" on:click={nextMonth}>
          <span class="sr-only">Right Arrow</span>
          <span class="next-month-arrow">
              <FaAngleRight/>
          </span>
        </button>
      </div>
    <div class="calendar-years">
      <button class="icon-button" on:click={previousYear}>
          <span class="sr-only">Left Arrow</span>
          <span class="prev-year-arrow"><FaAngleLeft/></span>
      </button>
      <button class="year-button {$selectionMode === 'year' ? 'active' : ''}
      " on:click={() => toggleSelectionMode('year')}>{currentYear}</button>
      <button class="icon-button" on:click={nextYear}>
          <span class="sr-only">Right Arrow</span>
          <span class="next-year-arrow"><FaAngleRight/></span>
      </button>
    </div>
  </div>
<div class="the-calendar">
  <div class="calendar-days">
      <p>Sun</p><p>Mon</p><p>Tus</p><p>Wed</p><p>Thu</p><p>Fri</p><p>Sat</p>
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
    <div class="dates-below">

    <button class="range-button {$selectionMode === 'range' ? 'active' : ''}" on:click={() => toggleSelectionMode('range')}>Range</button>
    <button class="clear-button" on:click={() => clearSelection()}>Clear</button>
    </div>  
  </div>
  
 
  <style>
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
    .calendar-months, .calendar-years {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 10px;
      padding-left: 10px;
      padding-right: 10px;
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


    .month-button, .year-button {
        background-color: transparent;
        border: none;
        color: black;
        cursor: pointer; 
        font-size: 1rem; 
        transition: transform 0.1s ease; 
        }

    .month-button:hover, .year-button:hover {
        transform: bold;
    }

    .month-button.active, .year-button.active {
        background-color: var(--primary-button-color);
        color: var(--primary-button-text-color);
        padding: 5px;
        border-radius: 10px;
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

    .day.today {
      border: var(--component-border)
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

    .calendar-controls {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 10px;
        padding-top: 10px;
        border-radius: 10px;
    }



    .range-button, .clear-button {
        background-color: var(--primary-button-color);
        color: white;
        padding: 5px;
        border-radius: 5px;
        cursor: pointer;
        border: 0;
    }

    .range-button:hover, .clear-button:hover {
        background-color: #68a1f1;
    }

    .range-button.active {
        background-color: var(--primary-button-hover-color);
        color: var(--primary-button-text-color);
    }

    .sr-only {
        position: absolute;
        width: 1px;
        height: 1px;
        padding: 0;
        margin: -1px;
        overflow: hidden;
        clip: rect(0, 0, 0, 0);
        white-space: nowrap;
        border: 0;
    }

  </style>