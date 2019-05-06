clear all;
close all;
clc;

data = csvread('meanfield_isingdata_CSV.csv', 1);

Temperature = data(:,1);
Magnetisation = data(:,2);
EnergyPerSite = data(:,3);
SpecificHeat = data(:,4);
MagneticSusceptibility = data(:,5);


plot(Temperature, Magnetisation, 'x-')

hold off;
xlabel('$\frac{\beta}{\beta_c}$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$\mathcal{M}$', 'Interpreter', 'latex', 'FontSize', 16);
