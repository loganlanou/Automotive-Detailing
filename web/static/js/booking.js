class BookingApp {
	constructor(root) {
		this.root = root;
		this.availabilityEndpoint = root.dataset.availabilityEndpoint || '/api/bookings/availability';
		this.submitEndpoint = root.dataset.submitEndpoint || '/api/bookings';
		this.daysContainer = root.querySelector('[data-calendar-days]');
		this.rangeLabel = root.querySelector('[data-calendar-range]');
		this.slotContainer = root.querySelector('[data-slot-list]');
		this.selectionPill = root.querySelector('[data-selection-pill]');
		this.form = document.getElementById('booking-form');
		this.feedback = document.getElementById('booking-feedback');
		this.navButtons = root.querySelectorAll('[data-month-nav]');
		this.state = {
			days: [],
			range: null,
			selectedDate: null,
			selectedSlotId: null,
			selectedSlotWindow: '',
			selectedSlotLabel: '',
		};

		this.bindNav();
		this.bindForm();
		this.loadAvailability();
	}

	bindNav() {
		this.navButtons.forEach((btn) => {
			btn.addEventListener('click', () => {
				const direction = btn.dataset.monthNav;
				if (direction === 'next') {
					const nextStart = this.nextStartDate();
					if (nextStart) this.loadAvailability(nextStart);
				} else {
					const prevStart = this.previousStartDate();
					if (prevStart) this.loadAvailability(prevStart);
				}
			});
		});
	}

	bindForm() {
		if (!this.form) return;
		this.form.addEventListener('submit', async (event) => {
			event.preventDefault();
			if (!this.state.selectedDate || !this.state.selectedSlotId) {
				this.showFeedback('Select a date and time before sending your request.', true);
				return;
			}

			const formData = new FormData(this.form);
			const payload = {
				name: (formData.get('name') || '').trim(),
				email: (formData.get('email') || '').trim(),
				phone: (formData.get('phone') || '').trim(),
				vehicle: (formData.get('vehicle') || '').trim(),
				service: (formData.get('service') || '').trim(),
				notes: (formData.get('notes') || '').trim(),
				date: this.state.selectedDate,
				slot_id: this.state.selectedSlotId,
			};

			this.setSubmitting(true);
			try {
				const response = await fetch(this.submitEndpoint, {
					method: 'POST',
					headers: { 'Content-Type': 'application/json' },
					body: JSON.stringify(payload),
				});
				const data = await response.json();
				if (!response.ok) {
					throw new Error(data.error || 'Unable to submit booking right now.');
				}
				this.showFeedback(data.message || 'Request received!', false);
				this.form.reset();
				this.clearSelection();
				this.loadAvailability(this.state.range ? this.state.range.start : undefined);
			} catch (error) {
				this.showFeedback(error.message || 'Unable to submit booking right now.', true);
			} finally {
				this.setSubmitting(false);
			}
		});
	}

	async loadAvailability(startDate) {
		if (this.rangeLabel) {
			this.rangeLabel.textContent = 'Refreshing…';
		}

		const params = new URLSearchParams({ days: 35 });
		if (startDate) {
			params.set('start', startDate);
		}

		try {
			const response = await fetch(`${this.availabilityEndpoint}?${params.toString()}`);
			if (!response.ok) {
				throw new Error('Unable to load availability');
			}
			const data = await response.json();
			this.state.days = data.days || [];
			this.state.range = data.range;
			if (this.state.selectedDate && !this.state.days.some((day) => day.date === this.state.selectedDate)) {
				this.state.selectedDate = null;
				this.state.selectedSlotId = null;
				this.state.selectedSlotLabel = '';
				this.state.selectedSlotWindow = '';
			}
			this.renderCalendar();
			this.renderSlots(this.state.selectedDate);
			this.updateNavStates();
			if (this.rangeLabel && data.range) {
				this.rangeLabel.textContent = this.formatRangeLabel(data.range);
			}
		} catch (error) {
			if (this.daysContainer) {
				this.daysContainer.innerHTML = `<div class="col-span-full text-center text-sm text-muted py-6">${error.message}</div>`;
			}
		}
	}

	renderCalendar() {
		if (!this.daysContainer) return;
		if (!this.state.days.length) {
			this.daysContainer.innerHTML = '<div class="col-span-full text-center text-sm text-muted py-6">No calendar data available.</div>';
			return;
		}

		const today = this.today();
		const isMobile = window.innerWidth < 640;

		this.daysContainer.innerHTML = this.state.days
			.map((day) => {
				const dateNumber = parseInt(day.date.split('-')[2], 10);
				const topLabel = day.label.split(',')[0] || '';
				const openSlots = day.slots ? day.slots.filter((slot) => slot.available).length : 0;
				const statusText = day.is_closed
					? 'Closed'
					: day.has_availability
						? `${openSlots} open`
						: 'Full';
				const isSelected = day.date === this.state.selectedDate;
				const disabled = day.is_closed || !day.has_availability;
				const highlight = day.is_today ? 'text-brand-accent' : 'text-muted';

				// Mobile-optimized compact classes
				const baseClasses = [
					'rounded-lg',
					'sm:rounded-xl',
					'border',
					'border-border',
					'p-2',
					'sm:p-3',
					'text-center',
					'sm:text-left',
					'transition',
					'focus-visible:outline-none',
					'focus-visible:ring-2',
					'focus-visible:ring-brand-accent',
					'active:scale-95',
				];

				if (isSelected) {
					baseClasses.push('border-brand-accent', 'bg-brand-accent/10', 'shadow-md', 'ring-1', 'ring-brand-accent/50');
				} else if (disabled) {
					baseClasses.push('opacity-40', 'cursor-not-allowed');
				} else {
					baseClasses.push('hover:border-brand-accent', 'cursor-pointer', 'hover:bg-white/5');
				}

				// Status indicator color
				const statusColor = day.is_closed
					? 'text-slate-500'
					: day.has_availability
						? 'text-emerald-400'
						: 'text-rose-400';

				const todayBadge = day.is_today
					? '<span class="absolute -top-1 -right-1 w-2 h-2 bg-brand-accent rounded-full sm:hidden"></span><span class="hidden sm:inline text-[9px] sm:text-[10px] uppercase tracking-wide text-brand-accent font-semibold">Today</span>'
					: '';
				const fullBadge = !day.is_closed && !day.has_availability
					? '<span class="hidden sm:inline text-[9px] sm:text-[10px] uppercase tracking-wide text-rose-400 font-semibold">Full</span>'
					: '';

				return `
					<button
						type="button"
						class="${baseClasses.join(' ')} relative"
						data-date="${day.date}"
						${disabled ? 'disabled' : ''}
					>
						${todayBadge.includes('absolute') ? todayBadge.split('<span class="hidden')[0] : ''}
						<div class="hidden sm:flex items-center justify-between mb-1">
							<span class="text-[10px] sm:text-xs uppercase tracking-wide ${highlight}">${topLabel}</span>
							${todayBadge.includes('hidden sm:inline') ? todayBadge.split('</span>')[1] + '</span>' : ''}${fullBadge}
						</div>
						<div class="sm:hidden text-[10px] uppercase tracking-wide ${highlight} mb-0.5">${topLabel}</div>
						<p class="text-xl sm:text-2xl font-heading font-bold ${day.is_today ? 'text-brand-accent' : ''}">${dateNumber}</p>
						<p class="text-[10px] sm:text-xs ${statusColor} mt-0.5 sm:mt-1 font-medium">${statusText}</p>
					</button>
				`;
			})
			.join('');

		this.daysContainer.querySelectorAll('button[data-date]').forEach((btn) => {
			btn.addEventListener('click', () => {
				this.handleDateSelect(btn.dataset.date);
			});
		});
	}

	renderSlots(date) {
		if (!this.slotContainer) return;
		if (!date) {
			this.slotContainer.innerHTML = '<div class="border border-border rounded-xl p-4 text-muted text-sm col-span-full">Select a date to see available times.</div>';
			return;
		}

		const day = this.state.days.find((d) => d.date === date);
		if (!day || day.is_closed || !day.slots.length) {
			this.slotContainer.innerHTML = '<div class="border border-border rounded-xl p-4 text-muted text-sm col-span-full">No sessions available for this day.</div>';
			return;
		}

		this.slotContainer.innerHTML = day.slots
			.map((slot) => {
				const disabled = !slot.available;
				const isSelected = slot.id === this.state.selectedSlotId && date === this.state.selectedDate;
				const baseClasses = [
					'rounded-xl',
					'border',
					'border-border',
					'p-4',
					'text-left',
					'transition',
					'flex',
					'flex-col',
					'gap-1',
				];

				if (isSelected) {
					baseClasses.push('border-brand-accent', 'bg-brand-accent/10', 'shadow-md');
				} else if (disabled) {
					baseClasses.push('opacity-50', 'cursor-not-allowed');
				} else {
					baseClasses.push('hover:border-brand-accent', 'cursor-pointer');
				}

				return `
					<button
						type="button"
						class="${baseClasses.join(' ')}"
						data-slot-id="${slot.id}"
						data-date="${date}"
						${disabled ? 'disabled' : ''}
					>
						<span class="text-sm uppercase tracking-wide text-muted">${slot.label}</span>
						<span class="text-lg font-heading font-semibold text-brand-fg">${slot.window}</span>
						${disabled ? '<span class="text-xs text-muted">Reserved</span>' : ''}
					</button>
				`;
			})
			.join('');

		this.slotContainer.querySelectorAll('button[data-slot-id]').forEach((btn) => {
			btn.addEventListener('click', () => {
				const slotId = btn.dataset.slotId;
				this.handleSlotSelect(date, slotId);
			});
		});
	}

	handleDateSelect(date) {
		if (!date) return;
		if (this.state.selectedDate !== date) {
			this.state.selectedSlotId = null;
			this.state.selectedSlotLabel = '';
			this.state.selectedSlotWindow = '';
		}
		this.state.selectedDate = date;
		this.renderCalendar();
		this.renderSlots(date);
		this.updateFormSelections();
	}

	handleSlotSelect(date, slotId) {
		const day = this.state.days.find((d) => d.date === date);
		if (!day) return;
		const slot = day.slots.find((s) => s.id === slotId);
		if (!slot || !slot.available) return;

		this.state.selectedDate = date;
		this.state.selectedSlotId = slotId;
		this.state.selectedSlotLabel = slot.label;
		this.state.selectedSlotWindow = slot.window;
		this.renderSlots(date);
		this.renderCalendar();
		this.updateFormSelections();
	}

	updateFormSelections() {
		if (!this.form) return;
		const dateInput = this.form.querySelector('input[name="selected_date"]');
		const slotInput = this.form.querySelector('input[name="slot_id"]');
		if (dateInput) dateInput.value = this.state.selectedDate || '';
		if (slotInput) slotInput.value = this.state.selectedSlotId || '';

		if (this.selectionPill) {
			if (this.state.selectedDate && this.state.selectedSlotId) {
				this.selectionPill.textContent = `${this.formatHumanDate(this.state.selectedDate)} • ${this.state.selectedSlotWindow}`;
				this.selectionPill.classList.remove('hidden');
			} else if (this.state.selectedDate) {
				this.selectionPill.textContent = `Great — now choose a slot on ${this.formatHumanDate(this.state.selectedDate)}.`;
				this.selectionPill.classList.remove('hidden');
			} else {
				this.selectionPill.classList.add('hidden');
			}
		}
	}

	clearSelection() {
		this.state.selectedDate = null;
		this.state.selectedSlotId = null;
		this.state.selectedSlotLabel = '';
		this.state.selectedSlotWindow = '';
		this.renderCalendar();
		this.renderSlots();
		this.updateFormSelections();
	}

	nextStartDate() {
		if (!this.state.range) return null;
		return this.addDays(this.state.range.end, 1);
	}

	previousStartDate() {
		if (!this.state.range) return null;
		const prev = this.addDays(this.state.range.start, -(this.state.range.days || 30));
		const today = this.today();
		if (this.compareDates(prev, today) < 0) {
			return this.state.range.start === today ? null : today;
		}
		if (prev === this.state.range.start) {
			return null;
		}
		return prev;
	}

	updateNavStates() {
		const today = this.today();
		this.navButtons.forEach((btn) => {
			if (btn.dataset.monthNav === 'prev') {
				const disabled = !this.state.range || this.compareDates(this.state.range.start, today) <= 0;
				btn.disabled = disabled;
				btn.classList.toggle('opacity-40', disabled);
			}
		});
	}

	formatRangeLabel(range) {
		const start = this.formatHumanDate(range.start);
		const end = this.formatHumanDate(range.end);
		return `${start} → ${end}`;
	}

	formatHumanDate(dateStr) {
		const date = new Date(`${dateStr}T12:00:00`);
		return date.toLocaleDateString(undefined, { weekday: 'short', month: 'short', day: 'numeric' });
	}

	addDays(dateStr, days) {
		const date = new Date(`${dateStr}T12:00:00`);
		date.setDate(date.getDate() + days);
		return date.toISOString().slice(0, 10);
	}

	compareDates(a, b) {
		if (a === b) return 0;
		return a > b ? 1 : -1;
	}

	today() {
		return new Date().toISOString().slice(0, 10);
	}

	setSubmitting(isSubmitting) {
		const button = this.form ? this.form.querySelector('button[type="submit"]') : null;
		if (!button) return;
		button.disabled = isSubmitting;
		button.classList.toggle('opacity-60', isSubmitting);
		button.querySelector('span')?.classList.toggle('animate-pulse', isSubmitting);
	}

	showFeedback(message, isError) {
		if (!this.feedback) return;
		this.feedback.textContent = message;
		this.feedback.classList.remove('hidden');
		this.feedback.classList.remove('border-rose-400/60', 'bg-rose-500/10', 'text-rose-100', 'border-brand-accent/60', 'bg-brand-accent/10', 'text-brand-fg');
		if (isError) {
			this.feedback.classList.add('border-rose-400/60', 'bg-rose-500/10', 'text-rose-100');
		} else {
			this.feedback.classList.add('border-brand-accent/60', 'bg-brand-accent/10', 'text-brand-fg');
		}
		setTimeout(() => {
			this.feedback.classList.add('hidden');
		}, 6000);
	}
}

document.addEventListener('DOMContentLoaded', () => {
	const root = document.getElementById('booking-app');
	if (root) {
		new BookingApp(root);
	}
});
