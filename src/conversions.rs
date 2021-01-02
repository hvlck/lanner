pub enum System {
    Metric,
    Imperial,
}

#[derive(Debug, Clone)]
pub enum Unit {
    LengthSystem,
    TemperatureSystem,
    Duration,
}

pub mod time {
    pub enum Duration {
        Millisecond,
        Second,
        Minute,
        Hour,
    }
}

pub mod mass {}

pub mod length {
    pub enum ImperialLengthUnit {
        Miles,
        Yards,
        Feet,
        Inches,
    }
    pub enum MetricLengthUnit {
        Kilometre,
        Hectometre,
        Decametre,
        Metre,
        Decimetre,
        Centimetre,
        Millimetre,
    }

    pub enum LengthSystem {
        ImperialLengthUnit,
        MetricLengthUnit,
    }
}

pub mod temperature {
    pub enum TemperatureSystem {
        Fahrenheit,
        Celsius,
        Kelvin,
    }
}

pub mod Currency {
    
}
