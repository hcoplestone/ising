clear all;
close all;
clc;

system = 0;
betaLowerLimit = 25;
betaUpperLimit = 100;
betaStep = 10;

figure;
hold on;
for beta = betaLowerLimit:betaStep:betaUpperLimit
    fname = ['./section2/beta-', num2str(beta) ,'-system', num2str(system) ,'.csv'];
    data = csvread(fname, 1);

    Sweep = data(:,1);
    Beta = data(:,2);
    SubSystemID = data(:,3);
    Magnetisation = data(:,4);
    DimensionlessEnergy = data(:,5);
    
    plot(Sweep(1:100), Magnetisation(1:100), 'x', 'DisplayName', sprintf('$\\beta = %0.2f$', beta/100))
end

hold off;
xlabel('Number of sweeps')
ylabel('$\mathcal{M}$', 'Interpreter', 'latex', 'FontSize', 16);

legend_handle = legend('-DynamicLegend');
set(legend_handle,'Interpreter','latex')
legend('show');

 
% figure;
% plot(Sweep, DimensionlessEnergy)
% ylabel('$E/J$', 'Interpreter', 'latex', 'FontSize', 16);