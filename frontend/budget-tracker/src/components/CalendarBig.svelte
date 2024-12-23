<script>
    import FaAngleLeft from 'svelte-icons/fa/FaAngleLeft.svelte'
    import FaAngleRight from 'svelte-icons/fa/FaAngleRight.svelte'
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';
    import { selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate, selectedDay } from '../stores/selectionMode';
    import { expenses } from '../stores/expenses';
    

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

    function formatDate(date) {
        const d = new Date(date);
        return `${d.getFullYear()}-${d.getMonth() + 1}-${d.getDate()}`;
    }

    // Convert selectedDate to the simplified format for comparison
    $: formattedSelectedDate = formatDate(selectedDate);

    function toDisplayDate(date) {
        const locale = "nb-NO";

        const year = date.getFullYear()
        const monthName = date.toLocaleString(locale, { month: "long" });
        const day = date.getDate();
        return `${day} ${monthName} ${year}`;
    }

  </script>
  
  <style>
    .calendar-box {
        display: flex;
        flex-direction: row;
        width: 100%;
        height: 100%;
        max-width: 50rem;
        max-height: 40rem;
        margin-top: 40px;


        background-color: var(--component-bg-color);
        border: 1px solid var(--component-border-color);
        border-radius: var(--component-border-radius);
        box-shadow: var(--component-box-shadow);
        font-family: var(--font-family);
    }

    .day-info {
        display: flex;
        flex-direction: column;
        align-items: left;
        width: 70%;
        margin: 1rem;
        padding: 0rem;
    }

    .date {
        font-size: 1.5rem;

    }
    
    
    .expenses-list {
        list-style-type: none;
        padding: 0;

    }

    .expense-item {
        display: flex;
        flex-direction: column;
        padding: 5px;
        border-bottom: 1px solid #ccc;
    }

    .expense-name {
        font-size: 1rem;
        font-weight: bold;
    }

    .expense-amount {
        font-size: 1rem;
    }

    .day-name {
        font-size: 1.5rem;
        margin-bottom: 1rem;
        font-weight: bold;
        text-transform: uppercase;
        font-family: var(--font-family);
    }
    .display-date {
        display: flex;
        flex-direction: column;
        align-items: left;
        margin-bottom: 1rem;
        text-transform: capitalize;
        margin-top: 2rem;
    }
    .calendar-full {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
        max-width: 50rem;
        max-height: 50rem;
        padding: 20px;

        background-color: var(--component-bg-color);

        border-radius: var(--component-border-radius);
        box-shadow: var(--component-box-shadow);
        font-family: var(--font-family);
    }

    .calendar-container {
        display: grid;
        grid-template-columns: repeat(7, 1fr);
        grid-gap: 4px;
        padding-top: 5px;
        background-color: var(--component-bg-color);      
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
        text-align: left;
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
        color: white;
        cursor: pointer; 
        font-size: 1rem; 
        transition: transform 0.1s ease; 
        }

    .month-button:hover, .year-button:hover {
        transform: scale(1.1);
        transform: bold;
    }

    .month-button.active, .year-button.active {
        background-color: var(--primary-button-color);
        color: var(--primary-button-text-color);
    }

    .previous-days {
        color: #ccc;
    }
    
    .day {
        display: flex;
        align-items: left;
        justify-content: space-between;
        height: 4rem;
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
        color: white;
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
        background-color: var(--primary-button-color);
        color: var(--primary-button-text-color);
        padding-top: 10px;
        border-radius: 10px;
    }



    .clear-button, .add-expense {
        background-color: var(--primary-button-color);
        color: white;
        padding: 0.7rem;
        border-radius: 5px;
        cursor: pointer;
        border: 0;
    }

    .clear-button:hover, .add-expense:hover {
        background-color: #68a1f1;
    }

    .day-content {
        display: flex;
        width: 100%;
        flex-direction: row;
        justify-content: space-between;
    }

    .number {
        font-size: 1rem;
    }

    .day-expense-number {
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 0.5rem;
        height: 0.7rem;
        width: 0.7rem;
        background-color: var(--primary-button-color);

        color: white;
        padding: 0.1rem;
        padding-right: 0;
        margin-left: 0.5rem;
        border-radius: 50%;
    }

    .day-expense-number.none {
        display: none;
    }

  </style>
  
  <div class="calendar-box">
    <div class="day-info">
        <div class="display-date">
            <span class="day-name">{selectedDate.toLocaleString("nb-NO", { weekday: "long" })},</span>
            <span class="date">{toDisplayDate(selectedDate)}</span>
        </div>
        <div>
        {#if $expenses.filter(exp => formatDate(exp.Date) === formattedSelectedDate).length === 0}
            <p>No expenses.</p>
        {:else if $expenses.filter(exp => formatDate(exp.Date) === formattedSelectedDate).length === 1}
            <p>1 expense</p>
            <ul class="expenses-list">
                {#each $expenses.filter(exp => formatDate(exp.Date) === formattedSelectedDate) as expense}
                    <li class="expense-item">
                        <span class="expense-name">{expense.Name}</span>
                        <span class="expense-amount">{expense.Amount} kr</span>
                    </li>
                {/each}
            </ul>
        {:else}
            <p>{$expenses.filter(exp => formatDate(exp.Date) === formattedSelectedDate).length} expenses</p>
            <ul class="expenses-list">
                {#each $expenses.filter(exp => formatDate(exp.Date) === formattedSelectedDate) as expense}
                    <li class="expense-item">
                        <span class="expense-name">{expense.Name}</span>
                        <span class="expense-amount">{expense.Amount} kr</span>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>
        <button class="add-expense">Add expense</button>
    </div>
    <div class="calendar-full">
    <div class="calendar-controls">
      <div class="calendar-months">
        <button class="icon-button" on:click={previousMonth}>
            <span class="prev-month-arrow"><FaAngleLeft/></span>
        </button>
        <button class="month-button {$selectionMode === 'month' ? 'active' : ''}
        " on:click={() => toggleSelectionMode('month')}>{
            (new Date(currentYear, currentMonth).toLocaleString('default', { month: 'long' })
              .replace(/^\w/, (c) => c.toUpperCase()))
          }</button>
      
        <button class="icon-button" on:click={nextMonth}>
          <span class="next-month-arrow">
              <FaAngleRight/>
          </span>
        </button>
      </div>
    <div class="calendar-years">
      <button class="icon-button" on:click={previousYear}>
          <span class="prev-year-arrow"><FaAngleLeft/></span>
      </button>
      <button class="year-button {$selectionMode === 'year' ? 'active' : ''}
      " on:click={() => toggleSelectionMode('year')}>{currentYear}</button>
      <button class="icon-button" on:click={nextYear}>
          <span class="next-year-arrow"><FaAngleRight/></span>
      </button>
    </div>
  </div>
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
          <div class="day-content">
          <span class="number">{date ? date.getDate() : ''}</span>
          
            {#if $expenses.filter(exp => formatDate(exp.Date) === formatDate(date)).length === 0}
            <span class="day-expense-number none">0</span>
                {:else}
                <span class="day-expense-number">{$expenses.filter(exp => formatDate(exp.Date) === formatDate(date)).length}</span>
                {/if}
            
        </div>
      </button>
  {/each}
    </div>
  </div>
    <div class="dates-below">
    <button class="clear-button" on:click={() => clearSelection()}>Today</button>
    </div>  
</div>
  </div>
  
 
  