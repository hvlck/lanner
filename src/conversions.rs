// std

// crates
use uom::fmt::DisplayStyle::Abbreviation;
use uom::si::f64::*;
use uom::si::length::{centimeter, kilometer, meter};
use uom::si::{Quantity, Units};
use uom::{impl_from, Conversion};

use crate::AstNode;

// local

//pub fn convert(node: AstNode) -> AstNode {
//    match node.as_rule() {};
//}

// This seems horribly inefficient but I couldn't find a generic conversion method in crate::uom
#[derive(Debug, Clone, Copy)]
pub enum LannerUnit {
    Angstrom,
    AstronomicalUnit,
    Attometer,
    Centimeter,
    Chain,
    Decameter,
    Decimeter,
    Exameter,
    Fathom,
    Femtometer,
    Fermi,
    Foot,
    FootSurvey,
    Gigameter,
    Hectometer,
    Inch,
    Kilometer,
    LightYear,
    Megameter,
    Meter,
    Micrometer,
    Micron,
    Mil,
    Mile,
    MileSurvey,
    Millimeter,
    Nanometer,
    NauticalMile,
    Parsec,
    Petameter,
    PicaComputer,
    PicaPrinters,
    Picometer,
    PointComputer,
    PointPrinters,
    Rod,
    Terameter,
    Yard,
    Yoctometer,
    Yottameter,
    Zeptometer,
    Zettameter,
}
